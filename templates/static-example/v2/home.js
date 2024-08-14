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
    cols[ColumnSize - 1].innerHTML = '<button class="btn" onclick="modifyRow(this)">修改</button>\n<button class="btn" onclick="removeRow(this)">删除</button>'
    // 为很一个单元格指定类名
    for (let i = 0; i < ColumnSize; ++i) 
        cols[i].className = "data";
}

// 删除行
async function removeRow(button) {
    try {
    let row = button.parentNode.parentNode;
    // 获取第一个单元格
    let id = row.querySelector('td').textContent;
    let res = await fetch('http://192.168.18.128:8888/del?id=' + id, {
        method:'DELETE'
    });
    console.log(res);
    if (!res.ok) {
        alert('error');
        return;
    }
    alert('operate ok')
    row.parentNode.removeChild(row);
    } catch(err) {
        console.log(err);
    }
}

// 获取整行的数据，发出patch请求
async function modifyRow(button) {
    let row = button.parentNode.parentNode;

    let jsonObj = {
        id: Number(row.cells[0].textContent),
        name: row.cells[1].textContent,
        gender: row.cells[2].textContent,
        introduction: row.cells[3].textContent,
        password: row.cells[4].textContent,
    };

    console.log(JSON.stringify(jsonObj));

    let res = await fetch('http://192.168.18.128:8888/mod', {
        method: 'PATCH',
        headers: {
            'Content-Type':'application/json'
        },
        body: JSON.stringify(jsonObj)
    });

    if (res.ok) {
        alert("operate ok!");
    } else {
        alert('occur error');
    }
}

// 内联编辑
document.addEventListener('DOMContentLoaded', () => {
    // 获取表格示例
    const table = document.getElementsByTagName('table')[0];

    table.addEventListener('click', function(event) {
        // 触发事件的单元格
        const target = event.target;
        // 处理可编辑的单元格
        if (target.classList.contains('editable')) {
            const currentContent = target.innerText;
            // 创建一个输入框
            const input = document.createElement('input');
            input.type = 'text';
            input.value = currentContent;
            input.classList.add('editing');
            
            target.innerHTML = '';
            target.appendChild(input);

            input.focus();

            // 为输入框绑定内容修改的事件
            input.addEventListener('blur', () => {
                const newContent = input.value.trim();
                if (newContent !== currentContent) {
                    target.innerText = newContent;
                    // TODO: fetch请求
                    console.log('updata value', newContent);
                } else {
                    target.innerText = currentContent;
                }
            });
        }
    }); // () =>
}); // document.addEventListener()