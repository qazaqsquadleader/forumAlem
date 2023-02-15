const baseURL = 'https://your-api-endpoint.com';

export const getData = () => {
  return fetch(`${baseURL}/data`, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json'
    }
  })
    .then(response => response.json())
    .catch(error => console.error(error));
};

export const postData = (data) => {
  return fetch(`${baseURL}/data`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(data)
  })
    .then(response => response.json())
    .catch(error => console.error(error));
};

/*              IMPLEMENTATION IN COMPONENTS
import React, { useState, useEffect } from 'react';
import { getData, postData } from './api';

const Component = () => {
  const [data, setData] = useState([]);

  useEffect(() => {
    const fetchData = async () => {
      const response = await getData();
      setData(response);
    };
    fetchData();
  }, []);

  const handleSubmit = async (e) => {
    e.preventDefault();
    const response = await postData({ data: 'some data' });
    console.log(response);
  };

  return (
    <div>
      <form onSubmit={handleSubmit}>
        <button type='submit'>Submit</button>
      </form>
      {data.map((item) => (
        <p key={item.id}>{item.name}</p>
      ))}
    </div>
  );
};

export default Component;


                AND RENDER CONDITIONALLY IN COMPONENTS

const ExampleComponent = () => {
const [data, setData] = useState(null);

return <div>Loading...</div>;
}

return <div>{data}</div>;
};

export default ExampleComponent;
        

*/