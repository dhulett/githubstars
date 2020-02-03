<template>
  <div>
    <SearchBar defaultQuery="" v-on:search-submitted="githubQuery" />
    <table>
      <th>
        <td v-for="header in tableHeader" v-bind:key="header">{{header}}</td>
      </th>
      <tr v-for="repo in repositories" v-bind:key="repo.ID">
        <td><a href="repo.URL">{{repo.Name}}</a></td>
        <td>{{repo.Description}}</td>
        <td>{{getLanguanges(repo.Languages)}}</td>
        <td>{{getTagsDisplay(repo.Tags)}}</td>
        <td>edit</td>
      </tr>
    </table>
  </div>
</template>

<script>
import SearchBar from './SearchBar.vue'
import githubClient from '../githubClient'
import { mapMutations, mapGetters, mapActions } from 'vuex'

export default {
  name: 'Home',
  components: { SearchBar },
  data() {
    return {
      tableHeader: ['Repository', 'Description', 'Language', 'Tags', ''],
      repositories: []
    }
  },
  computed: mapGetters(['allTags', 'repos']),
  created() {
    this.getTags();
  },
  methods: {
    githubQuery(query) {
      githubClient
        .getJSONRepos(query)
        .then(response => this.resetRepos(response.items) )
    },
    ...mapMutations(['resetRepos']),
    ...mapActions(['getTags']),
    selectedUser() {

    }
  },
}
</script>

<style>
</style>
