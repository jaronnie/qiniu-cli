<template>
  <div class="big" style="height: 100%">
    <el-row class="tac" style="height: 100%">
      <el-col :span="12" style="height: 100%">
        <el-menu
          default-active="this.$router.path" router mode ="horizontal"
          class="el-menu-vertical-demo"
          @open="handleOpen"
          @close="handleClose" style="height: 100%"
          background-color="#545c64"
          text-color="#fff"
          active-text-color="#ffd04b">
        
<!--          <el-submenu index="1">-->
<!--            <template slot="title">-->
<!--              <i class="el-icon-location"></i>-->
<!--              <span>element</span>-->
<!--            </template>-->
<!--            <el-menu-item-group>-->
<!--              <el-menu-item index="/element" key="1">选项1</el-menu-item>-->
<!--              <el-menu-item index="1-2">选项2</el-menu-item>-->
<!--            </el-menu-item-group>-->
<!--          </el-submenu>-->
<!--          <el-submenu index="2">-->
<!--            <template slot="title">-->
<!--              <i class="el-icon-menu"></i>-->
<!--              <span>Echart</span>-->
<!--            </template>-->
<!--            <el-menu-item-group>-->
<!--              <el-menu-item index="/Echart" key="2">选项1</el-menu-item>-->
<!--              <el-menu-item index="1-2">选项2</el-menu-item>-->
<!--            </el-menu-item-group>-->
<!--          </el-submenu>-->
          <el-menu-item index="/upload" key="1">
            <i class="el-icon-menu"></i>
            <span slot="title">upload</span>
          </el-menu-item>

        </el-menu>
      </el-col>
    </el-row>
  </div>
</template>

<script>
export default {
  name: "Dashboard",
  data() {
    return {
      navList: [
        {name:"/upload", navItem:"upload"},
      ],
      isTabBar: false
    }
  },
  methods: {
    handleOpen(key, keyPath) {
      console.log(key, keyPath);
    },
    handleClose(key, keyPath) {
      console.log(key, keyPath);
    }
  }
}
</script>

<style scoped>
 .big {
   position: fixed;
   top: 0;
   left: 0;
   width: 410px;
   margin-left: 0;
   height: 100%;
 }
</style>
