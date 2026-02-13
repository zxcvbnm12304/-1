<?php
require_once __DIR__ . '/vendor/autoload.php';

use phpspider\core\db;
use phpspider\core\log;
use phpspider\core\phpspider;
use phpspider\core\queue;
use phpspider\core\requests;
use phpspider\core\selector; 
// 第一步：数据库连接代码（完整）
$dbHost = 'localhost';
$dbUser = 'root';
$dbPwd = '123456';
$dbName = 'test';
$dbCharset = 'utf8mb4';
$dbPort = 3306;

try {
    $dsn = "mysql:host={$dbHost};dbname={$dbName};port={$dbPort};charset={$dbCharset}";
    // $pdo = new PDO($dsn, $dbUser, $dbPwd);
    // $pdo->setAttribute(PDO::ATTR_ERRMODE, PDO::ERRMODE_EXCEPTION);
    // $pdo->setAttribute(PDO::ATTR_EMULATE_PREPARES, false);
    echo "数据库连接成功！<br>";

    // 第二步：直接写新增代码（复用上面的$pdo实例）
    // $sql = "INSERT INTO ceshi1 (id,name,age) VALUES (:id,:name,:age)";
    // $stmt = $pdo->prepare($sql);
    // $params = [
    //     ":id" => 1,
    //     ':name'  => '张三',
    //     ':age'   => 21,
    // ];
    // $stmt->execute($params);

    // $affectedRows = $stmt->rowCount();
    // $lastInsertId = $pdo->lastInsertId();
    // if ($affectedRows > 0) {
    //     echo "新增成功！受影响行数：{$affectedRows}，新增ID：{$lastInsertId}";
    // } else {
    //     echo "新增失败，无数据插入";
    // }
} catch (PDOException $e) {
    // 连接+新增的错误都能被捕获
    die("操作出错：" . $e->getMessage());
}
?>