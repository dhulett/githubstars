import Vue from "vue";
import VueRouter from "vue-router";

import SelectUser from "./components/SelectUser";
import UserRepos from "./components/UserRepos";

Vue.use(VueRouter)

export default new VueRouter({
	mode: "history",
	routes: [{
			name: "home",
			path: "/",
			component: SelectUser
		},
		{
			name: "user",
			path: "/:user",
			component: UserRepos
		}
	]
})
