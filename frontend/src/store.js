import Vue from 'vue';
import Vuex from 'vuex';
import githubClient from './githubClient.js'

Vue.use(Vuex);

const store = new Vuex.Store({
    state: {
        repos: {},
        tags: {}
    },
    mutations: {
        resetRepos(state, repos) {
            state.repos = repos;
        },
        resetTags(state, kudos) {
            state.kudos = kudos;
        }
    },
    getters: {
        allTags(state) {
            return Object.values(state.tags);
        },
        tags(state) {
            return state.kudos;
        },
        repos(state) {
            return state.repos;
        }
    },
    actions: {
        async getRepos({
            commit
        }) {
            const starredRepos = await githubClient.getUserStarredRepos()

            commit('resetRepos', starredRepos)
        },
        // setTags({ commit, state }, repoId, tags) {
        //     commit('setTags', { repoId, })
        // },
        // addTag({ commit, state }, repoId, tag) {
        //     kudo => commit('resetKudos', { [kudo.id]: kudo, ...state.kudos })
        // },
        // removeTag({ commit, state }, repoId, tag) {

        //     const kudos = Object.entries(state.kudos).reduce((acc, [repoId, kudo]) => {
        //         return (repoId == repo.id) ? acc
        //             : { [repoId]: kudo, ...acc };
        //     }, {})

        //     commit('resetKudos', kudos)
        // }
    }
});

export default store;
