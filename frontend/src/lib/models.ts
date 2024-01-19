export interface AnimeSeason {
    season: string;
    year: number;
    animes: Anime[];
}
  
export interface License {
    name: string;
    url: string;
}
  
export interface Anime {
    index: number;
    sources: string[];
    title: string;
    type: string;
    episodes: number;
    status: string;
    animeSeason: AnimeSeason;
    picture: string;
    thumbnail: string;
    synonyms: string[];
    relations: string[];
    tags: string[];
}
  
export interface AnimeData {
    license: License;
    repository: string;
    lastUpdate: string;
    data: Anime[];
}
export interface Data<T>{
    msg : String;
    status : number;
    data : T;
}