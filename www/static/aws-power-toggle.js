//
// javascript functionality for AWS Power Toggle
// source: https://github.com/gbolo/aws-power-toggle
//


// listen for clicks to start buttons
$(document).on('click','.start-button',function(){
    var button = $(this)[0];
    button.disabled = true;

    $.ajax({
        type: "POST",
        dataType: "json",
        url: '/api/v1/env/' + $(this).data('env') + '/start',

        beforeSend: function(){
            button.innerHTML = '<i class="fas fa-hourglass-start"></i> <strong>sending</strong>';
        },
        success: function(stopResponse) {
            button.innerHTML = '<i class="fas fa-check"></i> <strong>success</strong>';
        },

        error: function(errorResponse) {
            button.innerHTML = '<i class="fas fa-times"></i> <strong>failed</strong>';
        },
    });
});

// listen for clicks on stop buttons
$(document).on('click','.stop-button',function(){
    var button = $(this)[0];
    button.disabled = true;

    $.ajax({
        type: "POST",
        dataType: "json",
        url: '/api/v1/env/' + $(this).data('env') + '/stop',

        beforeSend: function(){
            button.innerHTML = '<i class="fas fa-hourglass-start"></i> <strong>sending</strong>';
        },
        success: function(stopResponse) {
            button.innerHTML = '<i class="fas fa-check"></i> <strong>success</strong>';
        },

        error: function(errorResponse) {
            button.innerHTML = '<i class="fas fa-times"></i> <strong>failed</strong>';
        },
    });
});


// refresh button
$('#refreshButton').click(function() {
    writeEnvs()
});

// appends cards to env row
function drawEnvs(envDataAjax) {
    // target dom
    var rowEnv = document.getElementById('envRow');

    // loop through all envs
    for(var i = 0; i < envDataAjax.length; i++) {
        var env = envDataAjax[i];
        var running = env.running_instances + '/' + env.total_instances

        iconStart = document.createElement('i');
        iconStart.classList.add("fa", "fa-play");
        iconStop = document.createElement('i');
        iconStop.classList.add("fa", "fa-stop");
        labelRunning = document.createElement('span');
        labelRunning.classList.add("badge", "badge-success");
        labelRunning.innerText = "running";
        labelStopped = document.createElement('span');
        labelStopped.classList.add("badge", "badge-danger");
        labelStopped.innerText = "stopped";
        labelChanging = document.createElement('span');
        labelChanging.classList.add("badge", "badge-warning");
        labelChanging.innerText = "changing";
        labelMixed = document.createElement('span');
        labelMixed.classList.add("badge", "badge-warning");
        labelMixed.innerText = "mixed";

        buttonStart = document.createElement('button');
        buttonStart.classList.add("btn", "btn-info", "start-button");
        buttonStartText = document.createElement('strong');
        buttonStartText.innerText = " START";
        buttonStart.appendChild(iconStart);
        buttonStart.appendChild(buttonStartText);
        buttonStart.dataset.env = env.id;

        buttonStop = document.createElement('button');
        buttonStop.classList.add("btn", "btn-secondary", "stop-button");
        buttonStopText = document.createElement('strong');
        buttonStopText.innerText = " STOP";
        buttonStop.appendChild(iconStop);
        buttonStop.appendChild(buttonStopText);
        buttonStop.dataset.env = env.id;

        divCardBody = document.createElement('div');
        divCardBody.classList.add("card-body");
        divCardBody.innerHTML = '<p class="card-text">'+ running +' running instance(s)</p>';
        divCardBody.appendChild(buttonStart);
        divCardBody.appendChild(buttonStop);

        divCardHeader = document.createElement('h5');
        divCardHeader.classList.add("card-header");
        divCardHeader.textContent = env.name;
        if (env.state == "running") {
            divCardHeader.appendChild(labelRunning);
        } else if (env.state == "stopped") {
            divCardHeader.appendChild(labelStopped);
        } else if (env.state == "changing") {
            divCardHeader.appendChild(labelChanging);
        } else {
            divCardHeader.appendChild(labelMixed);
        }

        divCard = document.createElement('div');
        divCard.classList.add("card");
        divCard.appendChild(divCardHeader);
        divCard.appendChild(divCardBody);

        divCol = document.createElement('div');
        divCol.classList.add("col-env", "col-12", "col-sm-6", "col-lg-4");
        divCol.appendChild(divCard);

        // append it to dom
        rowEnv.appendChild(divCol);
    }
}

// gets the env data and calls drawEnvs on success
async function writeEnvs() {
    // show loading screen
    $('#envRow').html('<div class="col-12" id="loading"><div class="progress" style="height: 40px;"><div class="progress-bar progress-bar-striped progress-bar-animated bg-info" role="progressbar" aria-valuenow="100" aria-valuemin="0" aria-valuemax="100" style="width: 100%"><strong>Loading...</strong></div></div></div>');

    // get the data via API call
    await sleep(500);
    $.ajax({
        type: "GET",
        dataType: "json",
        url: '/api/v1/env/summary',

        success: function(envDataResponse) {
            // clean loading screen
            $('#envRow').empty();
            // modify the dom on the page
            drawEnvs(envDataResponse);
        },

        error: function(errorResponse) {
            // clean loading screen
            $('#envRow').empty();
            // display an alert
            errorHtml = `
                <div class="col-12" id="error">
                    <div class="alert alert-warning" role="alert">
                        <i class="fas fa-exclamation-triangle"></i>
                        There was an error with the request. Please try again or check server logs...
                        <hr /><i class="fas fa-asterisk"></i>
                        `+ errorResponse.responseText +`
                    </div>
                </div>`;
            $('#envRow').html(errorHtml);
        },
    });
}

// generic sleep function
function sleep(ms) {
    return new Promise(resolve => setTimeout(resolve, ms));
}

// draw the env on first load
writeEnvs();