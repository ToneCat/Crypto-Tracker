import React, { Component } from 'react';
import MUIDataTable from "mui-datatables";
import './App.css';

const columns = ["Currency", "Price"];

const options = {
  serverSide: true,
  onTableChange: (action, tableState) => {
    this.xhrRequest('http://localhost:3001/prices', result => {
      this.setState({ data: result });
    });
  }
};

<MUIDataTable 
  title={"Employee List"} 
  data={data} 
  columns={columns} 
  options={options} 
/>