import http from 'k6/http';
import {check} from 'k6';

/**
 * The K6 options. Below I have added some options that is worth highlighting.
 *
 * @link https://k6.io/docs/using-k6/options/#using-options
 */
export let options = {
    // We can configure k6 to discard the response bodies. This will decrease the load on the machine.
    // But only possible if we don't read the body later.
    //
    // @link https://k6.io/docs/using-k6/options/#discard-response-bodies
    discardResponseBodies: false,

    // We can restrict or remove the amount of redirections with 'maxRedirects'.
    // Set this value to 0 if you aren't expecting any.
    //
    // @link https://k6.io/docs/using-k6/options/#max-redirects
    maxRedirects: 0,

    // Thresholds can be a useful tool when we want to evaluate if a load test failed. A load test with
    // failed thresholds will be marked in Haubitze as 'Failed'. We can also use the thresholds to
    // abort the load test during execution. This load test is configured to abort if the rate
    // of successful checks goes below 60%, which indicates something is very wrong.
    //
    // @link https://k6.io/docs/using-k6/thresholds
    thresholds: {

        // Abort the test when the successful check rate goes below 60%.
        'checks': [
            {threshold: 'rate > 0.60', abortOnFail: true, delayAbortEval: '5s',}
        ],

        // Mark the test as 'Failed' if the average request duration is above 500ms or below 150ms.
        'http_req_duration': [
            'avg < 100',
        ],

    },

    // Scenarios defines the execution process of the load test. K6 will schedule VUs differently
    // depending on which executor is configured. We recommend the 'ramping-vus' if your load
    // test simulates a real user. 'ramping-arrival-rate' could prove to be more useful if
    // you are more interested in the RPS of a single endpoint.
    //
    // @link https://k6.io/docs/using-k6/scenarios/
    scenarios: {
        stress: {
            executor: 'ramping-arrival-rate',
            preAllocatedVUs: 50,
            maxVUs: 100,

            // Depending on your service, it's usually a good practice to ramp up the requests over a
            // period of time and then ramp down in the end. It does of course depend on how
            // your service normally receives traffic.
            stages: [
                {target: 2, duration: '5s'},
                {target: 0, duration: '5s'},
            ],
        },
    },

};


/**
 * K6(s) main loop.
 */
// eslint-disable-next-line no-undef
const baseUrl = __ENV.BASE_URL;

if (!baseUrl) {
    throw Error("BASE_URL is not defined");
}

const queryString = "userID=1";

export default function () {
    let res = http.get(`${baseUrl}/balance?${queryString}`);

    // You should always add check(s) in your load test. Check(s) will greatly increase the observability
    // of your load test and let you know if things are not running as expected.
    check(res, {
        'is status 200': (r) => r.status === 200,
    });

    // You can log a unexpected response. These logs should directly indicate that something is very wrong. 404(s) that
    // is somewhat expected shouldn't be logged as these will make the real errors harder to catch.
    if (res.status !== 200) {
        console.error(`Unexpected response: status=${res.status}, url=${res.request.url}, body=${res.body}`);
    }

}
