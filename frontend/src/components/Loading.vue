<template>
<div>
    <div></div>
    <div>{{loadingText}}</div>
</div>
</template>

<script>
export default {
    data() {
        return {
            loadingText: 'Getting the repositories list from Github...'
        }
    }
}
</script>

<style scoped>

</style>
