const VideoUploader = {
render() {
return <div class="video-uploader"> <div class="video-upload-area" onclick="VideoUploader.selectFile()"> <i class="fas fa-cloud-upload-alt"></i> <span>???? ????? ???? ????</span> <input type="file" id="videoInput" accept="video/*" style="display: none;"> </div> </div>;
},

selectFile() {
document.getElementById('videoInput').click();
}
};

window.VideoUploader = VideoUploader;


