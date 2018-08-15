<template>
  <div class="layout">
    <Layout>
      <Header :style="{position: 'fixed', width: '100%'}">
        <header-own></header-own>
      </Header>
      <Content :style="{margin: '88px 20px 0', background: '#fff', minHeight: '500px'}">
        <router-view/>
      </Content>
      <Footer class="layout-footer-center"></Footer>
    </Layout>
  </div>
</template>


<script>
  import headerOwn from './header.vue'
    export default {
        components: {
          headerOwn
        }
    }

</script>


<style scoped>
  .layout{
    border: 1px solid #d7dde4;
    background: #f5f7f9;
    position: relative;
    border-radius: 4px;
    overflow: hidden;
  }
  .layout-nav{
    width: 420px;
    margin: 0 auto;
    margin-right: 20px;
  }
  .layout-footer-center{
    text-align: center;
  }
</style>
