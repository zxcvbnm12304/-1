<?php
require_once __DIR__ . '/vendor/autoload.php';

use phpspider\core\db;
use phpspider\core\log;
use phpspider\core\phpspider;
use phpspider\core\queue;
use phpspider\core\requests;
use phpspider\core\selector; 

// 首页链接
$url1 = "https://b.faloo.com/1508571.html?1";
$html1 = requests::get($url1);
// var_dump($html1);
// exit;

// 目录页
$url2 ="https://b.faloo.com/html_1508_1508571/";
$html2 = requests::get($url2);
// var_dump($html2);exit;

// 目录页
//标题
$selector = "/html/body/div[2]/div[3]/div[2]/h1";
$title = selector::select($html2, $selector);
// var_dump($title);
// exit;

// 目录页
//作者
$selector ="/html/body/div[2]/div[3]/div[2]/div/a";
$author = selector::select($html2, $selector);
// var_dump($author);
// exit;

// 目录页
//时间
$selector =  "/html/body/div[2]/div[3]/div[2]/div/span[2]";
$time = selector::select($html2,$selector);
// var_dump($time);
// exit;

// 首页链接
//tupian
$selector = "/html/body/div[3]/div[2]/div[3]/div[2]/div[1]/a/img";
$img = selector::select($html1,$selector);
// var_dump($img);
// exit;

// 首页链接
//jianjie
$selector = "/html/body/div[3]/div[2]/div[3]/div[2]/div[2]/div[1]";
$Introduction = selector::select($html1,$selector);
// var_dump($Introduction);
// exit;

//目录
$selector = "/html/body/div[2]/div[3]/div[3]/div[3]//a/span";
$catalogue = selector::select($html2, $selector);
// var_dump($catalogue);
// exit;

//目录
//fufei
$selector = "/html/body/div[2]/div[3]/div[3]/div[6]/a/span";
$Paid = selector::select($html2, $selector);
// var_dump($Paid);
// exit;

//目录
//neirong
// $selector = "/html/body/div[2]/div[3]/div[4]/div[3]/div";
// $content = selector::select($html2, $selector);
// var_dump($content);
// exit;

$data = array(
    'title' => $title,
    'author' => $author,
    'time' => $time,
    'img' => $img,
    'Introduction' => $Introduction,
    'catalogue' => $catalogue,
    'Paid'=>$Paid,
);
// print_r($data);

$configs = array(
    'name' => '飞卢小说',
    'tasknum' => 8,
    'domains' => array(
        'b.faloo.com',
    ),
    'scan_urls' => array(
        'https://b.faloo.com/'
    ),
    'content_url_regexes' => array(
        "https://b.faloo.com/html_1508_1508571/"
    ),
    'list_url_regexes' => array(
        "https://b.faloo.com/1508571_1.html"
    ),
    'fields' => array(
        array(
            // 抽取内容页的文章免费章节
            'name' => "c_con_list",
            'selector' => "/html/body/div[2]/div[3]/div[3]/div[6]/a/span",
            'required' => true
        ),
        array(
            // 抽取内容页的文章neirong
            'name' => "noveContent readliner",
            'selector' => "/html/body/div[6]/div/div[1]/div[6]",
            'required' => true
        ),
    ),
   'db' => [
    'host' => '127.0.0.1',
    'user' => 'root',
    'pass' => '123456', 
    'name' => 'feilu',
    'tables' => 'fl',
    'charset' => 'utf8mb4',
    ], 
);
db::insert("content",$data);
$spider = new phpspider($configs);
$spider->start();
?>