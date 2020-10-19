import React from 'react';
import Alert from '@material-ui/lab/Alert';




class ScdlAlert extends React.Component {

    getAlert(){
    if(this.props.alert == "none"){

      return null
    }else if(this.props.alert == "good"){

      return <Alert variant="outlined" severity="success">Your song is downloading!</Alert>
    }else if(this.props.alert == "bad"){

      return <Alert variant="outlined" severity="error">Something is wrong! Check the URL</Alert>
    }
  }

    render() {
        return (
            <div>
              {this.getAlert()}
            </div>
          );
    }
  }






  




export default ScdlAlert;