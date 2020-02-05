const API_URL = "https://api.github.com/user/starred"

export default {
    authToken: "08c2fd2eb685f9c7fc456d69a7903a8ce02f4f49",
    async getUserStarredRepos() {
        const response = await fetch(`${API_URL}`, {
            headers: new Headers({
                'Authorization': `Bearer ${this.authToken}`
            })
        });
        return response.json();
    }
}
