<template>
  <div class="bg">
    <div class="upload">
      <h1>welcome to my free-pic</h1>
      <el-upload
        accept=".png, .jpg, .jpeg, .webp, .gif, mp3"
        drag
        action="http://upload-z2.qiniup.com"
        :before-upload="beforeUpload"
        :on-success="uploadSuccess"
        :data="upToken"
        multiple
        :limit="1">
        <i class="el-icon-upload"></i>
        <div class="el-upload__text">将文件拖到此处，或<em>点击上传</em></div>
      </el-upload>
      <el-button type="success" v-show="isSuccess" @click="copy(imgUrl)">点击复制链接</el-button>
    </div>
  </div>
</template>

<script>
export default {
  name: "page1",
  data() {
    return {
      imgName:"",
      imgUrl:"",
      isSuccess:false,
      upToken: {
        token: "",
        key: ""
      }
    }
  },
  methods: {
    beforeUpload(file) {
      let that = this;
      that.upToken.key = file.uid + "-" + file.name;
    },
    uploadSuccess(res) {
      let that = this;
      that.imgName = res.key;
      that.isSuccess = true;
      this.$notify({
        title:'upload successfully',
        type:"success"
      });
      that.imgUrl = "https://pic.gocloudcoder.com/" + this.imgName
      that.fileList.name = this.imgName
      that.fileList.url = this.imgUrl
    },
    getToken() {
      let that = this;
      this.$axios.get('http://gocloudcoder.com:8081/upload')
        .then(function (response) {
          console.log(response.data.token)
          that.upToken.token = response.data.token
        })
    },
    copy(data) {
      let url = data;
      let oInput = document.createElement('input');
      oInput.value = url;
      document.body.appendChild(oInput);
      oInput.select(); // 选择对象;
      console.log(oInput.value)
      document.execCommand("copy"); // 执行浏览器复制命令
      this.$message({
        message: '复制成功',
        type: 'success'
      });
      oInput.remove()
    }
  },
  mounted() {
    this.getToken()
  }
}
</script>

<style scoped>
.bg{
  text-align: center;
  background-color: #262d2c;
  border: 2px solid #1b1b23;
  padding-top: 0;
  height: calc(94vh - 62px);
}
.upload {
  padding-top: 100px;
}

</style>

