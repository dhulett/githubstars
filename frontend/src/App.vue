<template>
  <v-app>
    <div class="body">
      <Header applicationTitle="githubstars" homeButtonTitle="home" />
      <router-view />
    </div>
  </v-app>
</template>

<script>
import Header from "./components/Header";
export default {
  name: "App",
  components: {
    Header
  },
  data() {
    return {};
  }
};
</script>

<style>
body {
  font: 1em "Myriad Pro", Frutiger, "Lucida Grande", "Lucida Sans",
    "Lucida Sans Unicode", Verdana, sans-serif;
  margin: 0px;
  cursor: default;
}

.body {
  margin: 50px 50px;
}
</style>
