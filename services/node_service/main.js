var gmbh = require('gmbh');

var client;

function main(){

    // create new gmbh client
    client = new gmbh.gmbh();

    // assign service options as needed. Only a unique name is required, everything 
    // else can be left as default.
    client.opts.service.name = "nos";

    // Register routes
    client.Route("gatherData", handleOne);
    client.Route("gatherData", handleTwo);

    // start the client. The start method returns a promise and any follow up can be 
    // done using the .then()
    client.Start()
        .then(()=>{
            // things to be done after the client starts
        }).catch(()=>{
            // the onyl reason that the client will fail to start is if the name 
            // "CoreData" is chosen as the service name
        });
}

// handleOne
// Sender   String      the sender's name
// request  Payload     the payload data from the sender
function handleOne(sender, request){
    // use the client to create a new payload object
    let retval = client.NewPayload();

    // append the resut data to the payload
    retval.appendTextfields("result", `hello from nos, test 1; returning same message; message=${request.getTextfields('xid')}`)

    // all handler functions should return the payload that contains the results
    return retval;
}

// see handleOne above
function handleTwo(sender, request){
    let retval = client.NewPayload();
    retval.appendTextfields("result", `hello from nos, test 2; returning same message; message=${request.getTextfields('xid')}`)
    return retval;
}

main();