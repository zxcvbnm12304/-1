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

// //不按列表类型采集,以网址采集
// function get_all_list_pages1($base_url1) {
//     $page_urls = array($base_url1); // 初始页
//     // $page = 2; // 从第2页开始
    
//     // 循环检测分页是否存在，直到无法获取页面为止
//     while (true) {
//         $current_page_urls ="https://www.biquuge.com/1/{$page}/index_{$i}.html";
//          $html = requests::get($current_page_urls);
//          if(empty($html) || strpos($html, '404 Not Found') !== false){
//              break; // 页面不存在，停止分页
//          }
//          $page_urls[] = $current_page_urls;
//          for($page =1; $page <= 10;$page++,$i++){
//         //  $page++;
//         //  $i++;
//          if ($page +=1 && $i +=1){
//                     break;
//                     return; 
//                 }

//                 break;
//             }    
//         }
//     return $page_urls;
// }
// $all_list_pages1 = get_all_list_pages1("https://www.biquuge.com/1/1/index_1.html");
// echo "发现 " . count($all_list_pages1) . " 个分页需要采集\n";
// $all_novel_data = array();



// 列表采集,操作不顺
/**
 * 获取所有分页URL
 * @param string $base_url 基础URL
 * @return array 所有有效分页URL
 */
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
// 获取所有需要采集的分页
$all_list_pages = get_all_list_pages("https://www.biquuge.com/list1/");
echo "共发现 " . count($all_list_pages) . " 个分页需要采集\n";

// 存储所有采集到的小说数据
$all_novel_data = array();

foreach ($all_list_pages as $page_url) {
    echo "正在采集分页：{$page_url}\n";
    
    // 采集当前分页的HTML
    $html2 = requests::get($page_url);
    if (empty($html2)) {
        echo "分页 {$page_url} 抓取失败，跳过\n";
        continue;
    }
    // 类型
    $type = "/html/body/nav/div/a[2]";
    $type_name = selector::select($html2, $type);
    // 1. 采集小说标题
    $title_selector = "/html/body/section/div/div/div[1]/div/dl/dd[1]/h3/a";
    $novel_title = selector::select($html2, $title_selector);
    
    // 2. 采集小说链接（关键：获取a标签的href属性）
    // $link_selector = "/html/body/section/div/div/div[1]/div/dl/dd[1]/h3/a/@href";
    // $novel_link_relative = selector::select($html2, $link_selector);
    // $novel_link = BASE_DOMAIN . $novel_link_relative; // 拼接完整URL
    
    // 3. 采集作者
    $author_selector = "/html/body/section/div/div/div[1]/div/dl/dd[2]/span/a";
    $author = selector::select($html2, $author_selector);

    // 4. 采集状态
    $status_selector = "/html/body/section[1]/div/div/div[2]/div[2]/div/ul/li[4]";
    $status = selector::select($html2, $status_selector);

    // 5. 采集发布时间
    $time_selector = "/html/body/section/div/div/div[1]/div/dl/dd[4]";
    $create_time = selector::select($html2, $time_selector);

    // 6. 采集小说封面图片（拼接完整URL）
    $img_selector = "/html/body/section/div/div/div[1]/div/dl/dt/a/img/@src"; 
    $img_relative = selector::select($html, $img_selector);
    // $novel_cover = !empty($img_relative) ? BASE_DOMAIN . $img_relative : '';

    // 7. 采集小说简介
    // $intro_selector = "//*[@id='intro_pc']/text()[1]";
    // $novel_intro = selector::select($html2, $intro_selector);
    // $novel_intro = !empty($novel_intro) ? trim(preg_replace('/\s+/', ' ', $novel_intro)) : '';

    // 组装单本小说数据
   $data = array(
        // 'type_name'=>$type_name,
        // 'novel_title'=>$novel_title, 
        // 'novel_link' => $novel_link,
        // 'author'=> $author,
        'status'=> $status,
        // 'create_time'=>$create_time,
        // 'img_relative' => $img_relative,
        // 'novel_intro'=>$novel_intro,
        // 'crawl_time' => date('Y-m-d H:i:s'),
        // 'article'=>$novel_article,
        // 'article_link'=>$novel_article_links
    );
    print_r($data);
    print "----------\n";
    // 保存到总数据数组
    $all_novel_data[] = $data;
    
}

// echo "共采集到 " . count($all_novel_data) . " 本小说数据!\n";

// $dbHost = 'localhost';
// $dbUser = 'root';
// $dbPwd = '123456';
// $dbName = 'feilu2';
// $dbCharset = 'utf8mb4';
// $dbPort = 3306;

// try {
//     $dsn = "mysql:host={$dbHost};dbname={$dbName};port={$dbPort};charset={$dbCharset}";
//     $pdo = new PDO($dsn, $dbUser, $dbPwd);
//     $pdo->setAttribute(PDO::ATTR_ERRMODE, PDO::ERRMODE_EXCEPTION);
//     $pdo->setAttribute(PDO::ATTR_EMULATE_PREPARES, false);

//     $sql = "INSERT INTO biquuge (type_name,novel_title,author,create_time,status,Latest_chapter,novel_cover,Introduction_name,crawl_time)
//     VALUES (:type_name,:novel_title,:author,:create_time,:status,:Latest_chapter,:novel_cover,:Introduction_name,:crawl_time)";
//     $stmt = $pdo->prepare($sql);
//     $params = [
//         ':type_name'  => $type_name,
//         ':novel_title'  => $novel_title,
//         ':author'   => $author,
//         ':create_time'   => $create_time,
//         ':status'   => $status,
//         ':Latest_chapter'   => $Latest_chapter,
//         ':novel_cover'   => $novel_cover,
//         ':Introduction_name'   => $Introduction_name,
//         ':crawl_time'   => date('Y-m-d H:i:s')
//     ];
//     $stmt->execute($params);

//     $affectedRows = $stmt->rowCount();
//     $lastInsertId = $pdo->lastInsertId();
//     if ($affectedRows > 0) {
//         echo "新增成功！受影响行数：{$affectedRows},新增ID:{$lastInsertId}\n";
//     } else {
//         echo "新增失败，无数据插入\n";
//     }

// } catch (PDOException $e) {
//     die("操作出错：" . $e->getMessage() . "\n");
// }


$spider = new phpspider($configs);
$spider->start();
?>