import React, { Component } from 'react';
import { render } from 'react-dom';
import MUIDataTable from "mui-datatables";
import './App.css';


class App extends Component {

  // set the state (here we used states to ensure the API call was finished before rendering the table)
  state = {
    prices: []
  }

  pricesone = []; //setting array so we can fill it with JSON data from API Call
 
 
  
  componentDidMount() {
    fetch('http://localhost:3001/prices') // fetches the data from the go server
    .then(res => res.json())
    .then((res) => {
      for(var i in res) // converts JSON to array so it can be parsed and inserted with MUIDataTable
    this.pricesone.push([i, res [i].USD])
    this.setState({ prices: this.pricesone }) // sets current state to the array
    
      

    })
    .catch(console.log)
  }

  render(){
  const columns = ["Currency", "Price"]; // sets columns

  const options = {
    print: false,
    viewColumns : false,
    download: false,
    filter: false,
    selectableRows: false
  };

  return(
 



  <MUIDataTable 
    title={"Cryptocurrency Prices"} 
    data={this.state.prices} 
    columns={columns}
    options={options}/>
  
)
}
  
 
  
}
export default App;
render(<App />, document.getElementById('root'));