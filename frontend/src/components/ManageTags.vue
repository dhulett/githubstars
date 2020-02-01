<template>
  <v-container grid-list-md fluid class="grey lighten-4" >
    <v-layout align-center justify-space-around wrap>
      <v-flex md6>
        <h1 class="primary--text">
          <a :href="repo.html_url">{{repo.full_name}}</a>
        </h1>

        <v-chip class="text-xs-center">
          <v-avatar class="teal">
            <v-icon class="white--text">star</v-icon>
          </v-avatar>
          Stars: {{repo.stargazers_count}}
        </v-chip>

        <v-chip class="text-xs-center">
          <v-avatar class="teal white--text">L</v-avatar>
          Language: {{repo.language}}
        </v-chip>

        <v-chip class="text-xs-center">
          <v-avatar class="teal white--text">O</v-avatar>
          Open Issues: {{repo.open_issues_count}}
        </v-chip>
        <v-textarea
          name="input-7-1"
          label="Show some love"
          value=""
          v-model="repo.notes"
          hint="Describe why you love this project"
        ></v-textarea>
        <v-btn @click.prevent="updateKudo(repo)"> Kudo </v-btn>
        <router-link tag="a" to="/me">Back</router-link>
      </v-flex>
    </v-layout>
  </v-container>
</template>

<script>
import { mapActions, mapGetters } from 'vuex';

export default {
  data() {
    return {
      repo: {}
    }
  },
  watch: {
    '$route': 'fetchData'
  },
  computed: mapGetters(['kudos']),
  created() {
    this.fetchData();
  },
  methods: {
    async fetchData() {
      const response = await fetch('https://api.github.com/repositories/' + this.$route.params.id);
      const responseJson = await response.json();
      Object.assign(responseJson, this.kudos[this.$route.params.id])
    },
    ...mapActions(['updateKudo'])
  }
}
</script>

<style>
</style>
