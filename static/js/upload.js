document.addEventListener('DOMContentLoaded', function() {
    var uploadButton = document.getElementById('uploadButton');
    uploadButton.addEventListener('click', function() {
        var fileInput = document.getElementById('fileInput');
        var file = fileInput.files[0];
        var outputDiv = document.getElementById('output'); // 获取显示结果的元素

        if (!file) {
            alert('请选择一个文件进行上传。');
            return;
        }

        var formData = new FormData();
        formData.append('file', file);

        fetch('/r', {
            method: 'POST',
            body: formData
        })
        .then(response => {
            if (!response.ok) {
                throw new Error('网络响应不是ok状态！');
            }
            return response.text();
        })
        .then(data => {
            // 将响应的HTML显示在页面上
            outputDiv.innerHTML = data;
        })
        .catch(error => {
            console.error('上传文件失败：', error);
        });
    });
});
