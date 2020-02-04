<template>
<div>
    <div class="title"> {{ applicationTitle }} </div>
    <div class="home" v-if="displayHomeButton"><router-link to="/">{{ homeButtonTitle }}</router-link> </div>
</div>
</template>

<script>
export default {
    data() {
        return {
            displayHomeButton: false
        }
    },
    beforeRouteUpdate(to, from, next) {
        this.displayHomeButton = to == '/'
        next()
    },
    props: {
        applicationTitle: String,
        homeButtonTitle: String,
    }
}
</script>

<style scoped>
div {
    position: relative;
    display: flex;
}

.title {
    float: left;
}

.home {
    float: right;
}

</style>
