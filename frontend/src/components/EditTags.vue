<template>
 <div></div>
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
