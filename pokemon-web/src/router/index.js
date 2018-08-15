import Vue from 'vue'
import Router from 'vue-router'
import Layout from '@/components/common/layout'
import contentLayout from '@/components/common/content-layout'
import Prop from '@/components/prop/prop'

Vue.use(Router);

export default new Router({
  mode: 'history',
  routes: [
    {
      path: '/',
      name: 'pokemon',
      component: Layout,
      children: [
        {
          path: 'pokemon',
          name: '宝可梦',
          component: Prop,
          icon: "ios-paper",
        },
        {
          path: 'skill',
          name: '技能',
          component: Prop
        },
        {
          path: 'prop',
          name: '属性',
          component: Prop
        },
        {
          path: 'special',
          name: '特性',
          component: Prop
        },
        {
          path: 'type',
          name: '类型',
          component: Prop
        },
        {
          path: 'admin',
          name: '管理',
          component: contentLayout,
          children: [
            {
              path: 'pokemon',
              name: '管理宝可梦',
              component: Prop
            },
          ]
        },
      ]
    }
  ]
})
