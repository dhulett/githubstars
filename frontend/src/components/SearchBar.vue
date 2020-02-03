<template>
    <v-toolbar dark color="teal">
      <v-text-field
        solo-inverted
        flat
        hide-details
        label="search by tag"
        prepend-inner-icon="search"
        v-model="query"
        @keyup.enter="onSearchSubmition"
      ></v-text-field>
    </v-toolbar>
</template>

<script>
export default {
    data() {
      return {
        query: null,
      };
    },
    props: ['defaultQuery'],
    methods: {
      onSearchSubmition() {
        this.$emit('search-submitted', this.query);
      },
      async logout () {
        await this.$auth.logout()
        this.$router.push('/')
    }
  }
}
</script>
