const ColumnSize = 6;
const PromptInfomation = ['uid', '用户名', '性别', '简介', '密码', '操作'];

// 获取用户输入
function getUserInformation() 
{
    let data = new Array(ColumnSize - 1);
    for (let i = 0; i < ColumnSize - 1; ++i) {
        data[i] = prompt(PromptInfomation[i]);
        if (data[i] == null || data[i].trim() == "") {
            return null;
        }
    }
    return data;
}

function addRow()
{
    let table = document.getElementsByTagName('table')[0];
    let len = table.rows.length;

    let new_row = table.insertRow(len);
    let cols = new Array(ColumnSize);

    for (let i = 0; i < ColumnSize; ++i) {
        cols[i] = new_row.insertCell(i);
    }

    let data = getUserInformation();
    if (data == null) {
        return;
    }
    // 数据更新
    for (let i = 0; i < ColumnSize - 1; ++i) {
        cols[i].innerHTML = data[i];
    }
    cols[ColumnSize - 1].innerHTML = '<button class="mod_btn" onclick="modifyRow(this)">修改</button>\n<button class="del_btn" onclick="removeRow(this)">删除</button>'
    // 为很一个单元格指定类名
    for (let i = 0; i < ColumnSize; ++i) 
        cols[i].className = "data";
}

// 删除行
function removeRow(button) {
    let row = button.parentNode.parentNode;
    row.parentNode.removeChild(row);
}

// 编辑行
// TODO 编辑整行，应该更加精细
function modifyRow(button) {
    let row = button.parentNode.parentNode;
    console.log(row);

    let data = getUserInformation();
    if (data == null) {
        return;
    }

    for (let i = 0; i < ColumnSize - 1; ++i) {
        row.cells[i].innerText = data[i];
    }
}