<?php
$dbHost = 'localhost';
$dbUser = 'root';
$dbPwd = '123456';
$dbName = 'feilu2';
$dbCharset = 'utf8mb4';
$dbPort = 3306;

try {
    // 创建 PDO 连接
    $dsn = "mysql:host={$dbHost};dbname={$dbName};port={$dbPort};charset={$dbCharset}";
    $pdo = new PDO($dsn, $dbUser, $dbPwd);
    $pdo->setAttribute(PDO::ATTR_ERRMODE, PDO::ERRMODE_EXCEPTION);
    $pdo->setAttribute(PDO::ATTR_EMULATE_PREPARES, false);
    echo "数据库连接成功！<br>";

    // 构建 SQL 查询
    $sql = "SELECT id, novel_title, author, create_time, img, Introduction, catalogue, Paid, crawl_time FROM fl2";

    // 执行查询
    $stmt = $pdo->query($sql);

    // 检查结果集是否存在
    if ($stmt) {
        // 检查结果集中的行数
        $rowCount = $stmt->rowCount();

        if ($rowCount > 0) {
            // 输出数据
            while ($row = $stmt->fetch(PDO::FETCH_ASSOC)) {
                echo "id: " . $row["id"] . " - Name: " . $row["novel_title"] . " " . $row["author"] . 
                " " . $row["create_time"] . " " . $row["img"] ;
            }
        } else {
            echo "0 结果";
        }
    } else {
        echo "查询失败";
    }
} catch (PDOException $e) {
    die("数据库连接失败: " . $e->getMessage());
}
?>