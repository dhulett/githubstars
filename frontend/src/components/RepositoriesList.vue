<template>
  <div>
    <table>
      <tr>
        <th v-for="header in tableHeader" v-bind:key="header">{{ header }}</th>
      </tr>
      <tr v-for="repo in filterRepos(repositories)" v-bind:key="repo.id">
        <td>
          <a href="repo.URL">{{ repo.Name }}</a>
        </td>
        <td>{{ repo.description }}</td>
        <td>{{ repo.languages }}</td>
        <td>{{ getTagsDisplay(repo.tags) }}</td>
        <td onclick="this.$emit('editRepo', repo)">edit</td>
      </tr>
    </table>

    <EditTags v-if="showModal" @close="showModal = false" />
  </div>
</template>

<script>
export default {
  data() {
    return {
      showModal: true,
      repoInEdition: null,
      tableHeader: ["Repository", "Description", "Language", "Tags", ""]
    };
  },
  props: {
    repositories: Array
  },
  methods: {
    getLanguanges(languages) {
      return languages;
    },
    getTagsDisplay(tags) {
      return tags;
    },
    showEditTagsModal(repo) {
      this.repoInEdition = repo;
    },
    filterRepos(repos) {
      return repos;
    }
  }
};
</script>

<style scoped>
table {
  font-family: arial, sans-serif;
  border-collapse: collapse;
  width: 100%;
}

td,
th {
  font-size: 1em;
  border: 1px solid #dddddd;
  text-align: left;
  padding: 8px;
}

tr:nth-child(even) {
  background-color: #dddddd;
}
</style>
