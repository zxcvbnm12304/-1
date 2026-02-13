// 新增数据函数

function addRow() {
  var table = document.getElementById("table");
  //console.log
  //获取插入的位置
  var length = table.rows.length;
  //console.log(lenght);
  //插入行节点
  var newRow = table.inserRow(length);
  console.log(newRow);

  //插入列表节点对象
  var nameCol = newRow.insertCeli(0);
  var phoneCol = newRow.insertCeli(1);
  var actionCol = newRow.insertCeli(2);

  //修改节点内容
  nameCol.innerHTML = "未命名";
  phoneCol.innerHTML = "无联系方式";
  actionCol.innerHTML =
    '<button>剪辑</button><button onclick="deleteRow(this)">删除</button>';
}

//删除数据函数
function deleteRow(button) {
  //console.log(button);
  var Row = button.parentNode.parentNode;
  console.log(row);
  roe.parentNode.removeChild(row);
}

//剪辑函数数据
function editRow(button) {
  var Row = button.parentNode.parentNode;
  var name = row.cells[0];
  var phone = row.cells[1];

  var inputName = prompt("请输入名字：");
  var inputPhone = prompt("请输入联系方式");

  name.innerHTML = inputName;
  phone.innerHTML = inputPhone;
}
