<template>
  <div>
    <Menu mode="horizontal" theme="dark" :active-name="activeMenu">
      <template v-for="(route, $index) in $router.options.routes[$router.options.routes.length - 1].children">
        <template  v-if="route.children && route.name">
          <Submenu :name="route.name">
            <template slot="title">
              <Icon :type="route.icon" /> {{ route.name }}
            </template>
              <MenuItem
                v-for="(cRoute, $index) in route.children"
                :key="$index"
                :to="cRoute"
                :name="route.name">
                <Icon :type="cRoute.icon"></Icon>{{ cRoute.name }}
              </MenuItem>
          </Submenu>
        </template>

        <template v-else>
          <MenuItem :name="route.name" :to="route">
            <Icon :type="route.icon"></Icon>{{route.name}}
          </MenuItem>
        </template>

      </template>
    </Menu>
    <br>
  </div>
</template>
<script>
  export default {
    data () {
      return {
        activeMenu: ''
      }
    },
    watch: {
      $route(value){
        this.activeMenu = value.name;
      }
    },
    created(){
      this.activeMenu = this.$route.name;
    },
  }
</script>
