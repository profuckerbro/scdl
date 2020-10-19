import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Button from '@material-ui/core/Button';
import TextField from '@material-ui/core/TextField';
import Alert from '@material-ui/lab/Alert';
import ScdlAlert from './Alerts';
const axios = require('axios').default;



class ScdlDownload extends React.Component {

    constructor(props) {
        super(props);
        this.state = {url: '',alert: "none"};
    
        this.handleChange = this.handleChange.bind(this);
      }

    handleChange(event) {
       this.setState({url: event.target.value});
    }
    handleAlert(status){
      this.setState({alert: status})
    }



  PostRequest(url) {
     
    const payload = url
    console.log(payload)
axios({
    url: 'http://localhost:8080/download',
    method: 'POST',
    responseType: 'blob', // important
    data: {
        "Url": payload
        },
  }).then((response) => {
    console.log(response)
    const url = window.URL.createObjectURL(new Blob([response.data]));
    const link = document.createElement('a');
    link.href = url;
    //Set Download Name
    const contentDisposition = response.headers['content-disposition'];
    let fileName = 'unknown';
    if (contentDisposition) {
        const fileNameMatch = contentDisposition.match(/filename[^;=\n]*=(?:(\\?['"])(.*?)\1|(?:[^\s]+'.*?')?([^;\n]*))/i);
      
        fileName = fileNameMatch[2];
    }
  
    link.setAttribute('download', fileName);
    document.body.appendChild(link);
    link.click();
    this.handleAlert("good")
  }).catch((error) => this.handleAlert("bad"));


}



    render() {
        return (
            <div>
              <ScdlAlert alert={this.state.alert} />
            <form noValidate autoComplete="off" className={this.props.passName} onSubmit={() => this.PostRequest(this.state.url)} >
              <TextField id="outlined-basic" label="Soundcloud URL" variant="outlined" onChange={this.handleChange}/>
              <Button onClick={() => this.PostRequest(this.state.url)} >Download</Button>
            </form>
        
            </div>
          );
    }
  }






  




export default ScdlDownload;