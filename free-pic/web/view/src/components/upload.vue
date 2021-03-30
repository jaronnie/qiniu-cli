<template>
  <div class="bg">
    <el-upload
      accept=".png, .jpg"
      drag
      action="http://upload-z2.qiniup.com"
      :before-upload="beforeUpload"
      :on-success="uploadSuccess"
      :data="upToken"
      multiple>
      <i class="el-icon-upload"></i>
      <div class="el-upload__text">将文件拖到此处，或<em>点击上传</em></div>
    </el-upload>
    <el-button type="success" v-show="isSuccess" >链接</el-button>

    <img v-show="isSuccess" :src= this.imgUrl alt="" height="200" width="200">
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
      console.log(file)
      let that = this;
      that.upToken.key = file.uid + "-" + file.name;
    },
    uploadSuccess(res) {
      console.log(res)
      let that = this;
      that.imgName = res.key;
      that.isSuccess = true;
      alert("https://picture.nj-jay.com/"+this.imgName)
      that.imgUrl = "https://picture.nj-jay.com/" + this.imgName

    },
    getToken() {
      let that = this;
      this.$axios.get('http://localhost:8081/upload')
        .then(function (response) {
          console.log(response.data.token)
          that.upToken.token = response.data.token
        })
    }
  },
  mounted() {
    this.getToken()
  }
}
</script>

<style scoped>
</style>

