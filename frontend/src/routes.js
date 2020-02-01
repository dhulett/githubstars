import Vue from 'vue';
import VueRouter from 'vue-router';

import Home from './components/Home';
import ManageTags from './components/ManageTags';
import UserRepos from './components/UserRepos';

Vue.use(VueRouter);

export default new VueRouter({
    mode: 'history',
    routes: [
        {
            path: '/',
            component: Home
        },
        {
            path: '/user/:user',
            component: UserRepos
        },
        {
            path: '/user/:user/repo/:repo',
            component: ManageTags
        }
    ]
});
