<template>
  <div>
    <SearchBar v-on:searchSubmitted="filterTag($event)" />
    <Loading v-if="loading" />
    <RepositoriesList v-else v-bind:repositories="fetchData()" />
    <EditTags v-if="showTagEditingModal" />
  </div>
</template>

<script>
import SearchBar from "./SearchBar";
import RepositoriesList from "./RepositoriesList";
import Loading from "./Loading";
import EditTags from "./EditTags";
import githubClient from "../githubClient";

export default {
  components: { SearchBar, RepositoriesList, Loading, EditTags },
  data() {
    return {
      loading: true,
      showTagEditingModal: false
    };
  },
  created() {
    this.fetchData();
  },
  methods: {
    async fetchData() {
      this.loading = true;
      let repos = await githubClient.getUserStarredRepos();
      this.loading = false;
      return repos;
    }
  }
};
</script>

<style></style>
