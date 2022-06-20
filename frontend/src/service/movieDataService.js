import axios from 'axios'

const MOVIE_API_URL = "http://localhost:3000"

class movieDataService {

    getAllMovies() {
        return axios.get(`${MOVIE_API_URL}/movies`);
    }

    getIdMovie(id) {

        return axios.get(`${MOVIE_API_URL}/movie/${id}`);
    }

    deleteMovie(id) {

        return axios.delete(`${MOVIE_API_URL}/movie/${id}`);
    }

    updateMovie(id, movie) {

        return axios.put(`${MOVIE_API_URL}/movie/${id}`, movie);
    }

    createMovie(movie) {

        return axios.post(`${MOVIE_API_URL}/movie`, movie);
    }   
}

export default new movieDataService()