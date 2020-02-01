const API_URL = "https://api.github.com/search/repositories"

export default {
    async getJSONRepos(query) {
        const response = await fetch(`${API_URL}?q=` + query);
        return response.json();
    }
}
