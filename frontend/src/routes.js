import Vue from 'vue';
import VueRouter from 'vue-router';

import SelectUser from './components/SelectUser';
import EditTags from './components/EditTags';
import UserRepos from './components/UserRepos';

Vue.use(VueRouter);

export default new VueRouter({
    mode: 'history',
    routes: [
        {
            path: '/',
            component: SelectUser
        },
        {
            path: '/user/:user',
            component: UserRepos
        },
        {
            path: '/user/:user/repo/:repo',
            component: EditTags
        }
    ]
});
