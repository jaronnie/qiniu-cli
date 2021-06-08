<template>
  <div class="bg">
    <router-view></router-view>
  </div>
</template>

<script>
import Dashboard from "./view/dashboard/Dashboard";
export default {
  name: "App",
  data() {
    return {
      components: {
        Dashboard
      }
    }
  }
}
</script>

<style scoped>

</style>

