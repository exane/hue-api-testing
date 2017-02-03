var webPage = require('webpage');
var page = webPage.create();
var system = require('system');

var url = system.env.DASHBOARD_URL

page.open(url, function (status) {
  var url = page.url;

  var anything_failing = pipelinesFailing()
  // anything_failing ||= !AWSCostsFailing()
  // anything_failing ||= !WIPFailing()

  system.stdout.write(anything_failing);
  phantom.exit();
});

function pipelinesFailing() {
  return page.evaluate(function() {
    var failings = document.querySelectorAll(".widget-concourse-ci .failed").length;
    return failings > 0
  });
};
