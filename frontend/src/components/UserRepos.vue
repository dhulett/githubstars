<template>
  <div>
    <div >
      <SearchBar defaultQuery="selectedUser" v-on:search-submitted="githubQuery" />
      <RepositoriesList />
    </div>
    <div>
      <Loading />
    </div>
  </div>
</template>

<script>
import SearchBar from './SearchBar'
import RepositoriesList from './RepositoriesList'
import Loading from './Loading'
import githubClient from '../githubClient'
import { mapMutations, mapGetters, mapActions } from 'vuex'

export default {
  name: 'Home',
  components: { SearchBar, RepositoriesList, Loading },
  data() {
    return {
      showLoading: true
    }
  },
  computed: mapGetters(['allRepos', 'repos']),
  created() {
    this.getKudos();
  },
  methods: {
    githubQuery(query) {
      this.showLoading = true
      githubClient
        .getJSONRepos(query)
        .then(response => {
          this.showLoading = false
          this.resetRepos(response.items)
        })
    },
    ...mapMutations(['resetRepos']),
    ...mapActions(['getKudos']),
    selectedUser() {
      return this.user()
    }
  },
}
</script>

<style>
</style>
