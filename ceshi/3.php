<?php
require_once __DIR__ . '/vendor/autoload.php';

use phpspider\core\db;
use phpspider\core\log;
use phpspider\core\phpspider;
use phpspider\core\queue;
use phpspider\core\requests;
use phpspider\core\selector; 

// 新增：定义网站域名（用于拼接完整图片URL）
define('BASE_DOMAIN', 'https://www.biquuge.com');

// 新增：分页采集的核心配置（追加到你的configs里）
$configs = array(
    'name' => 'biquuge',
    'domains' => array(
        'biquuge.com',
    ),
    'scan_urls' => [
        'https://www.biquuge.com/list1/'  // 小说列表主页面
    ],
    'list_url_regexes' => [
        "https://www.biquuge.com/list1/\d+.html"
    ],
    'content_url_regexes' => [
        "https://www.biquuge.com/\d+/"  // 匹配小说详情页
    ],
    'auto_turn_page' => true,
    'interval' => 1,
    'retry' => 3,
    'fields' => array(
        array(
            'name' => "article_content",
            'selector' => "/html/body/section/div/div/div[1]/div/dl/dd[1]/h3/a",
            'required' => true,
        ),
        array(
            'name' => "article_author",
            'selector' => "/html/body/section/div/div/div[1]/div/dl/dd[2]/span/a",
            'required' => true
        ),
    ),
);

// 新增：分页采集的核心逻辑——循环获取所有分页的URL
function get_all_list_pages($base_url) {
    $page_urls = array($base_url); // 初始页
    $page = 2; // 从第2页开始
    
    // 循环检测分页是否存在，直到无法获取页面为止
    while (true) {
        $current_page_url = "https://www.biquuge.com/list1/{$page}.html";
        // 测试页面是否可访问
        $html = requests::get($current_page_url);
        if (empty($html) || strpos($html, '404 Not Found') !== false) {
            break; // 页面不存在，停止分页
        }
        $page_urls[] = $current_page_url;
        $page++;
        // 防止无限循环，设置最大分页数（可根据需要调整）
        if ($page > 100) {
            break;
        }
    }
    return $page_urls;
}

// 第一步：获取所有分页的URL（替代原有只抓第1页的逻辑）
$all_list_pages = get_all_list_pages("https://www.biquuge.com/list1/");
echo "共发现 " . count($all_list_pages) . " 个分页需要采集\n";

// 第二步：循环采集每个分页的小说数据
foreach ($all_list_pages as $page_url) {
    echo "正在采集分页：{$page_url}\n";
    
    // 采集当前分页的HTML
    $html2 = requests::get($page_url);
    if (empty($html2)) {
        echo "分页 {$page_url} 抓取失败，跳过\n";
        continue;
    }

    // 你的原有采集逻辑（仅修改图片URL处理）
    // 标题 
    $selector = "/html/body/section/div/div/div[1]/div/dl/dd[1]/h3/a";
    $novel_title = selector::select($html2,$selector);

    // 作者
    $selector ="/html/body/section/div/div/div[1]/div/dl/dd[2]/span/a";
    $author = selector::select($html2, $selector);

    // 状态
    $selector = "/html/body/section/div/div/div[1]/div/dl/dd[3]";
    $status = selector::select($html2,$selector);

    // 时间
    $selector =  "/html/body/section/div/div/div[1]/div/dl/dd[4]";
    $create_time = selector::select($html2,$selector);

    // 图片（关键修改：拼接完整URL，不保存本地）
    $selector = "/html/body/section/div/div/div[1]/div/dl/dt/a/img/@src"; // 新增@src获取图片路径
    $img_relative = selector::select($html2,$selector);
    $img = BASE_DOMAIN . $img_relative; // 拼接成完整URL（如：https://www.biquuge.com/images/129/129970/129970s.jpg）

    // 组装数据
    $data = array(
        'novel_title' => $novel_title,
        'author' => $author,
        'status' => $status, // 新增：补充状态字段
        'create_time' => $create_time,
        'img' => $img, // 现在是完整的图片URL
        'crawl_time' => date('Y-m-d H:i:s')
    );

    // 打印当前分页的采集结果
    print_r($data);
    print "-------------------------\n";
}

echo " 所有分页抓取数据完成！\n";

// 你的原有数据库配置（保留，未修改）
$dbHost = 'localhost';
$dbUser = 'root';
$dbPwd = '123456';
$dbName = 'feilu2';
$dbCharset = 'utf8mb4';
$dbPort = 3306;

// 如果你需要插入数据库，解除下面的注释（已适配新的img字段）

try {
    $dsn = "mysql:host={$dbHost};dbname={$dbName};port={$dbPort};charset={$dbCharset}";
    $pdo = new PDO($dsn, $dbUser, $dbPwd);
    $pdo->setAttribute(PDO::ATTR_ERRMODE, PDO::ERRMODE_EXCEPTION);
    $pdo->setAttribute(PDO::ATTR_EMULATE_PREPARES, false);

    $sql = "INSERT INTO biquuge (id,novel_title,author,status,create_time,img,crawl_time)
    VALUES (:id,:novel_title,:author,:status,:create_time,:img,:crawl_time)";
    $stmt = $pdo->prepare($sql);
    $params = [
        'id'=>1,
        ':novel_title'  => $novel_title,
        ':author'   => $author,
        ':status'   => $status,
        ':create_time'   => $create_time,
        ':img'   => $img, // 存储完整图片URL
        ':crawl_time'   => date('Y-m-d H:i:s')
    ];
    $stmt->execute($params);

    $affectedRows = $stmt->rowCount();
    $lastInsertId = $pdo->lastInsertId();
    if ($affectedRows > 0) {
        echo "新增成功！受影响行数：{$affectedRows}，新增ID：{$lastInsertId}\n";
    } else {
        echo "新增失败，无数据插入\n";
    }

} catch (PDOException $e) {
    die("操作出错：" . $e->getMessage() . "\n");
}


// 保留你原有爬虫启动逻辑（可选，如果你需要用phpspider的自动采集）
$spider = new phpspider($configs);
$spider->start();
?>