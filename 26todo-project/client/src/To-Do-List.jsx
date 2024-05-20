import React, {Component} from 'react';
import axios from 'axios';
import {Card, Header, Form, Icon, Input} from 'semantic-ui-react';

let endPoint="http://localhost:9000";

class ToDoList extends Component {
    constructor(props){
        super(props);
        this.state = {
            task: "",
            newTask: [],
        };
    }

    componentDidMount(){
        this.getTask();
    }

    onChange=(event)=>{
        this.setState({[event.target.name]:event.target.value})
    }

    onSubmit=(event)=>{
        let {task}=this.state;
        if(task){
            axios.post(endPoint+"/api/tasks",{task,},{headers:{"Content-Type":"application/x-www-form-urlencoded"},}).then((res)=>{
                this.getTask();
                this.setState({task:"",});
                console.log(res);
            });
        }
    };

    getTask=()=>{
        axios.post(endPoint+"/api/task").then((res)=>{
            if(res.data){
                this.setState({
                    items: res.data.map((item)=>{
                        let color="yellow";
                        let style={
                            wordWrap:"break-word",
                        };
                        if(item.status){
                            color="green";
                            style["textDecorationLine"]="line-through";
                        }
                        return (<Card key={item._id} color={color} fluid className="rough">
                            <Card.Content>
                                <Card.Header textAlign="left">
                                    <div style={style}>
                                        {item.task}
                                    </div>
                                </Card.Header>
                                <Card.Meta textAlign="right">
                                    <Icon name="check circle" color="blue" onClick={()=>this.updatetask(item._id)} />
                                    <span style={{paddingRight: 10}}>Undo</span>
                                    <Icon name="delete" color="red" onClick={()=>this.deleteTask(item._id)}/>
                                    <span style={{paddingRight: 10}}>Delete</span>
                                </Card.Meta>
                            </Card.Content>
                        </Card>);
                    }),
                });
            }
            else{
                this.setState({items:[],});
            }
        });
    };

    updatetask=(id)=>{
        axios.put(endPoint+"/api/task/"+id,{headers:{
            "Content-Type":"application/x-www-form-urlencoded",
        },}).then((res)=>{
            console.log(res);
            this.getTask();
        });
    }

    undoTask=(id)=>{
        axios.put(endPoint+"/api/undoTask/"+id,{headers:{
            "Content-Type":"application/x-www-form-urlencoded",
        },}).then((res)=>{
            console.log(res);
            this.getTask();
        });
    }

    deleteTask=(id)=>{
        axios.delete(endPoint+"/api/deleteTask/"+id,{headers:{
            "Content-Type":"application/x-www-form-urlencoded",
        },}).then((res)=>{
            console.log(res);
            this.getTask();
        });
    };

    render(){
        return(
            <div>
                <div className="row">
                    <header className="header" as="h2" color="yellow">
                        TO DO LIST
                    </header>
                    <div className="row">
                        <Form onSubmit={this.onSubmit}>
                            <input type="text" name="task" onChange={this.onChange} value={this.state.task} fluid placeholder="Create Task" />
                        </Form>
                        <div className="row">
                            <Card.Group>{this.state.items}</Card.Group>
                        </div>
                    </div>
                </div>
            </div>
        );
    }
}

export default ToDoList;