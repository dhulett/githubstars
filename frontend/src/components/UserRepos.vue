<template>
  <div>
    <Loading v-if="loading"/>
    <RepositoriesList v-else/>
  </div>
</template>

<script>
import RepositoriesList from './RepositoriesList'
import Loading from './Loading'
import githubClient from '../githubClient'
import { mapMutations, mapGetters, mapActions } from 'vuex'

export default {
  name: 'Home',
  components: { RepositoriesList, Loading },
  data() {
    return {
      loading: true
    }
  },
  computed: mapGetters(['allRepos', 'repos']),
  created() {
    this.getKudos();
  },
  methods: {
    async githubQuery(query) {
      this.loading = true
      let response = await githubClient.getJSONRepos(query)
      this.loading = false
      this.resetRepos(response.items)
    },
    ...mapMutations(['resetRepos']),
    ...mapActions(['getKudos']),
    selectedUser() {
      return this.user
    }
  },
}
</script>

<style>
</style>
