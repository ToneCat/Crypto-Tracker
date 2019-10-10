import React, { Component } from 'react';
import { render } from 'react-dom';
import MUIDataTable from "mui-datatables";
import './App.css';
import { createMuiTheme, MuiThemeProvider } from '@material-ui/core/styles';


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
    this.pricesone.push([i, res[i].USD])
    this.setState({ prices: this.pricesone }) // sets current state to the array
    })
    .catch(console.log)
  }

  render(){

    const columns = [
  
      {
        name: "Currency",
        options: {
          filter: true,
        }
      },
      {
        label: "Price (USD)",
        options: {
          filter: true,
        }
      },
      {
        name: "More Info",
        options: {
          filter: false,
          sort: false,
          empty: true,
          customBodyRender: (value, tableMeta, updateValue) => {
        
            return (
              <button onClick={() => 
            //  window.alert(`Clicked "Edit" for row ${tableMeta.rowIndex}`)
            window.open("https://cryptowat.ch/assets/"+tableMeta.rowData[0])
            

              }>
             
           
              <img id="chart" src="https://img.icons8.com/nolan/64/000000/positive-dynamic.png"></img>
            
              </button>
            );
          }
        }
        
      }
    ];

  const options = {
    print: false,
    viewColumns : false,
    download: false,
    filter: false,
    selectableRows: false,
    rowsPerPage:7
  };

  
 

  return(

    <div id="table">
  <MUIDataTable 
    title={"Cryptocurrency Prices"} 
    data={this.state.prices} 
    columns={columns}
    options={options}/>
    </div>
  
  
)
}

}

export default App;
render(<App />, document.getElementById('root'));