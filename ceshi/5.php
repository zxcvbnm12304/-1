<?php
require_once __DIR__ . '/vendor/autoload.php';

use phpspider\core\db;
use phpspider\core\log;
use phpspider\core\phpspider;
use phpspider\core\queue;
use phpspider\core\requests;
use phpspider\core\selector; 

// 定义基础域名，用于拼接完整URL
define('BASE_DOMAIN', 'https://www.biquuge.com');
requests::set_header('User-Agent', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36');
requests::set_header('Accept', 'text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8');
requests::set_header('Referer', BASE_DOMAIN);
requests::set_timeout(10);

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
        "https://www.biquuge.com/\d+/" , // 匹配小说详情页
        "https://www.biquuge.com/129/\id+/\d+.html",
        "https://www.biquuge.com/\d+/\d+/",
    ],
    'auto_turn_page' => true,
    'interval' => 1,
    'retry' => 3,
    'fields' => array(
        array(
            'name'=>"type_name",
            'selector'=>"/html/body/section[1]/div/div/div[1]/a[2]",
            'required' => true
        ),
        array(
            'name' => "article_title",
            'selector' => "/html/body/section/div/div/div[1]/div/dl/dd[1]/h3/a",
            'required' => true,
        ),
        array(
            'name' => "article_author",
            'selector' => "/html/body/section/div/div/div[1]/div/dl/dd[2]/span/a",
            'required' => true
        ),
        array(
            'name' => "article_link",
            'selector' => "/html/body/section/div/div/div[1]/div/dl/dd[1]/h3/a/@href",
            'required' => true
        ),
        array(
            'name' => "novel_cover",
            'selector' => "/html/body/section[1]/div/div/div[2]/div[1]/img/@src",
            'required' => true
        ),
        array(
            'name' => "Introduction_name",
            'selector' => "//*[@id='intro_pc']/br[1]",
            'required' => true
        ),
        array(
            'name' => "article",
            'selector'=> "/html/body/section[3]/div/div/div/div[1]/ul/li/a",
            'required' => true
        )
    ),
);

// 修复分页采集函数
function get_all_list_pages1($base_url1) {
    $page_urls = array($base_url1); // 初始页
    $page = 2; 
    
    while (true) {
        $current_page_urls = "https://www.biquuge.com/1/{$page}/index_1.html";
        $html = requests::get($current_page_urls);
        if(empty($html) || strpos($html, '404 Not Found') !== false || strlen($html) < 1000){
            break; 
        }
        $page_urls[] = $current_page_urls;
        $page++; 
        
        if ($page > 120) { 
            break;
        }
    }
    return $page_urls;
}

// ===================== 核心修改1：初始化数组缓存所有采集数据 =====================
$all_novel_data = array(); // 存储所有采集到的小说数据
$all_list_pages1 = get_all_list_pages1("https://www.biquuge.com/1/1/index_1.html");
echo "发现 " . count($all_list_pages1) . " 个分页需要采集\n";

foreach ($all_list_pages1 as $page_urls) {
    echo "正在采集分页：{$page_urls}\n";
    $html = requests::get($page_urls);
    if (empty($html)) {
        echo "分页 {$page_urls} 抓取失败，跳过\n";
        continue;
    }
    
    // 采集数据（保持原有选择器）
    $type = "/html/body/section[1]/div/div/div[1]/a[2]";
    $type_name = selector::select($html, $type);
    $title_selector = "/html/body/section[1]/div/div/div[2]/div[2]/h1";
    $novel_title = selector::select($html, $title_selector);
    $author_selector = "/html/body/section[1]/div/div/div[2]/div[2]/div/ul/li[1]/a";
    $author = selector::select($html, $author_selector);
    $time_selector = "/html/body/section[1]/div/div/div[2]/div[2]/div/ul/li[2]";
    $create_time = selector::select($html, $time_selector);
    $status_selector = "/html/body/section[1]/div/div/div[2]/div[2]/div/ul/li[4]";
    $status = selector::select($html, $status_selector);
    $Latest = "/html/body/section[1]/div/div/div[2]/div[2]/div/ul/li[5]/a";
    $latest_chapter = selector::select($html, $Latest);
    $img_selector = "/html/body/section[1]/div/div/div[2]/div[1]/img/@src";
    $img_relative = selector::select($html, $img_selector);
    $novel_cover = !empty($img_relative) ? BASE_DOMAIN . $img_relative : '';
    $Introduction = "//*[@id='intro_pc']/text()[4]";
    $Introduction_name = selector::select($html, $Introduction);
    
    // ===================== 核心修改2：清洗数据（避免数据库格式错误） =====================
    // 清洗时间（提取纯时间字符串）
    $timePattern = '/\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}/';
    preg_match($timePattern, $create_time, $timeMatches);
    $cleanCreateTime = $timeMatches[0] ?? date('Y-m-d H:i:s'); // 无有效时间则用当前时间
    
    // 清洗状态：只提取“：”后的文本，保留“连载/完结”
    $statusPattern = '/：([^；，。]+)/';
    preg_match($statusPattern, $status, $statusMatches);
    $statusText = trim($statusMatches[1] ?? $status); 
    // 统一状态文本格式
    $statusMap = ['连载' => '连载', '完结' => '完结', '完本' => '完结'];
    $cleanStatus = $statusMap[$statusText] ?? '连载';
    
    // ===================== 核心修复：排除id字段，不手动传入 =====================
    $data = array(
        'type_name' => $type_name ?: '', 
        'novel_title' => $novel_title ?: '',
        'author' => $author ?: '',
        'create_time' => $cleanCreateTime,
        'status' => $cleanStatus,
        'latest_chapter' => $latest_chapter ?: '',
        'novel_cover' => $novel_cover ?: '',
        'Introduction_name' => $Introduction_name ?: '',
        'data_time' => date('Y-m-d H:i:s'),
        // 绝对不要加id字段！数据库会自动自增生成
    );
    $all_novel_data[] = $data; 
    print_r($data);
    print "----------\n";
}

// ===================== 核心修改4：批量插入数据库（防重复+ID自增） =====================
$dbHost = 'localhost';
$dbUser = 'root';
$dbPwd = '123456';
$dbName = 'feilu2';
$dbCharset = 'utf8mb4';
$dbPort = 3306;

// 只有采集到数据才执行插入
if (!empty($all_novel_data)) {
    try {
        $dsn = "mysql:host={$dbHost};dbname={$dbName};port={$dbPort};charset={$dbCharset}";
        $pdo = new PDO($dsn, $dbUser, $dbPwd);
        $pdo->setAttribute(PDO::ATTR_ERRMODE, PDO::ERRMODE_EXCEPTION);
        $pdo->setAttribute(PDO::ATTR_EMULATE_PREPARES, false);

        // 开启事务
        $pdo->beginTransaction();
        
        // ========== 核心：明确指定插入字段（排除id），避免顺序混乱 ==========
        // 1. 手动指定字段列表（固定顺序，永不乱）
        $fields = [
            'type_name', 'novel_title', 'author', 'create_time', 
            'status', 'latest_chapter', 'novel_cover', 'Introduction_name', 'data_time'
        ];
        $fieldStr = implode(',', $fields);
        
        // 2. 构建更新语句（重复时更新，同样排除id）
        $updateStr = '';
        foreach ($fields as $field) {
            $updateStr .= "{$field} = VALUES({$field}),";
        }
        $updateStr = rtrim($updateStr, ','); 
        
        // 3. 构建批量插入的占位符和参数（按固定字段顺序）
        $sqlValues = [];
        $params = [];
        $index = 0;
        foreach ($all_novel_data as $item) {
            $placeholders = [];
            // 严格按预定义的fields顺序遍历，保证字段和值一一对应
            foreach ($fields as $key) {
                $placeholder = ":{$key}_{$index}";
                $placeholders[] = $placeholder;
                $params[$placeholder] = $item[$key]; // 按key取值，不会乱
            }
            $sqlValues[] = '(' . implode(',', $placeholders) . ')';
            $index++;
        }
        
        // 4. 最终SQL（无id字段，数据库自动自增）
        $sql = "INSERT INTO biquuge ({$fieldStr}) 
                VALUES " . implode(',', $sqlValues) . "
                ON DUPLICATE KEY UPDATE {$updateStr}";
        
        $stmt = $pdo->prepare($sql);
        $stmt->execute($params);
        
        $pdo->commit();
        
        // 统计结果（修正计数逻辑，更准确）
        $totalProcessed = count($all_novel_data);
        $affectedRows = $stmt->rowCount();
        $insertedCount = $affectedRows - $totalProcessed; // 插入返回1，更新返回2 → 差值为更新数
        $updatedCount = $totalProcessed - $insertedCount;
        echo "操作完成！新增：{$insertedCount} 条，更新：{$updatedCount} 条，总计处理：{$totalProcessed} 条\n";

    } catch (PDOException $e) {
        $pdo->rollBack();
        die("插入/更新出错：" . $e->getMessage() . "\n");
    }
} else {
    echo "未采集到任何小说数据，无需插入\n";
}

$spider = new phpspider($configs);
$spider->start();
?>