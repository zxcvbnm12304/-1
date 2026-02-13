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
/**
 * 获取所有分页URL
 * @param string $base_url 基础URL
 * @return array 所有有效分页URL
 */
function get_all_list_pages($base_url) {
    $page_urls = array($base_url); // 初始页
    // $page = 2; // 从第2页开始
    
    // 循环检测分页是否存在，直到无法获取页面为止
    while (true) {
        $current_page_url = "https://www.biquuge.com/1/{$page}/index_1.html";
        // 测试页面是否可访问
        $html = requests::get($current_page_url);
        if (empty($html) || strpos($html, '404 Not Found') !== false) {
            break; // 页面不存在，停止分页
        }
        $page_urls[] = $current_page_url;
        for($page = 1; $page <= 10; $page+=1){
        if ($page++) {
            break;
            return;
        }
        break;
        return;
       }
    }
    return $page_urls;
}
$all_list_pages = get_all_list_pages("https://www.biquuge.com/1/1/index_1.html");
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
    $type = "/html/body/section[1]/div/div/div[1]/a[2]";
    $type_name = selector::select($html2, $type);
    // 1. 采集小说标题
    $title_selector = "/html/body/section[1]/div/div/div[2]/div[2]/h1";
    $novel_title = selector::select($html2, $title_selector);
    // 3. 采集作者
    $author_selector = "/html/body/section[1]/div/div/div[2]/div[2]/div/ul/li[1]/a";
    $author = selector::select($html2, $author_selector);
    // 4. 采集发布时间
    $time_selector = "/html/body/section[1]/div/div/div[2]/div[2]/div/ul/li[2]";
    $create_time = selector::select($html2, $time_selector);
    // 5. 采集状态
    $status_selector = "/html/body/section[1]/div/div/div[2]/div[2]/div/ul/li[4]";
    $status = selector::select($html2, $status_selector);
    // 6.最新章节
    $Latest = "/html/body/section[1]/div/div/div[2]/div[2]/div/ul/li[5]/a";
    $Latest_chapter = selector::select($html2, $Latest);

    // 组装单本小说数据
    $novel_data = array(
        'type_name'=>$type_name,
        'novel_title'=>$novel_title, 
        'author'=>$author,
        'create_time'=>$create_time,
        'status'=>$status,
        'Latest_chapter'=>$Latest_chapter,
    );

    // 保存到总数据数组
    // $all_novel_data[] = $novel_data;
    
    // 打印当前小说的采集结果
    print_r($novel_data);
    print "-------------------------\n";
}