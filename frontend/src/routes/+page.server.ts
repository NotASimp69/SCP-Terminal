import { error } from '@sveltejs/kit';

/** @type {import('./$types').Actions} */
export const actions = {
    default: async ({request}) => {
      const data = await request.formData();
      const query = data.get('query');
      console.log(query);

      const response = await fetch('http://localhost:8080/fetch?query='+query, {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
        },
      });
      
      if(response.ok){
        const data = await response.json();
        return{
      data
        }
      }
      throw error(500,"Something went horribly wrong! :(");
    }
  };