<template>
  <div>
    <SearchBar v-on:search-submitted="githubQuery" />
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
      repositories: [
        {
          ID: "1",
          Name: "placeholderRepo1",
          URL: "https://github.com",
          Description: "Just one placeholder",
          Languages: ['lang1', 'lang2'],
          Tags: ['tag1', 'tag2']
        },
        {
          ID: "2",
          Name: "placeholderRepo1",
          URL: "https://github.com",
          Description: "Just one placeholder",
          Languages: ['lang1', 'lang2'],
          Tags: ['tag1', 'tag2']
        }
      ]
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

    },
    getLanguanges(languages) {
      return languages
    },
    getTagsDisplay(tags) {
      return tags
    }
  },
}
</script>

<style scoped>
table {
  font-family: arial, sans-serif;
  border-collapse: collapse;
  width: 100%;
}

td, th {
  border: 1px solid #dddddd;
  text-align: left;
  padding: 8px;
}

tr:nth-child(even) {
  background-color: #dddddd;
}
</style>
