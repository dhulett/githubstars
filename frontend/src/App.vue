<template>
 <v-app>
   <Header applicationTitle="githubstars" homeButtonTitle="home"/>
   <router-view/>
 </v-app>
</template>

<script>
import Header from "./components/Header"

export default {
  name: 'App',
  components: {
    Header
  },
  data() {
    return {}
  }
}
</script>

<style>

</style>
