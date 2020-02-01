<template>
 <v-app>
   <router-view></router-view>
 </v-app>
</template>

<script>
import Loading from './components/Loading'

export default {
 name: 'App',
 components: { Loading },
 data() {
   return {}
 }
}
</script>

<style>

</style>
