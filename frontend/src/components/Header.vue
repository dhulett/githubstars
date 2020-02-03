<template>
<div>
    <div> {{ applicationTitle }} </div>
    <div> {{ homeButtonTitle }} </div>
<div>
</template>

<script>
export default {
    data() {
        return {
            applicationTitle: "githubstars",
            homeButtonTitle: "home",
            displayHomeButton: false
        }
    }
}
</script>

<style scoped>

</style>
