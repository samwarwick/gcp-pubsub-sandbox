using Google.Api.Gax;
using Google.Cloud.PubSub.V1;

var projectId = "gps-demo";
var topicId = "demo-topic";

var publisherService = await new PublisherServiceApiClientBuilder
{
    EmulatorDetection = EmulatorDetection.EmulatorOrProduction
}.BuildAsync();

var topicName = new TopicName(projectId, topicId);

if (publisherService.GetTopic(topicName) == null)
{
    publisherService.CreateTopic(topicName);
}

var publisher = await PublisherClient.CreateAsync(
    topicName,
    new PublisherClient.ClientCreationSettings()
        .WithEmulatorDetection(EmulatorDetection.EmulatorOrProduction));

var subject = "Hello world!";
Console.WriteLine(args.Length);
if (args.Length > 0 && args[0] != "") {
    subject = args[0];
}

var json = $"{{ \"service\": \"dotnet\", \"subject\": \"{subject}\" }}";

Console.WriteLine($"SENDING: {json}");

await publisher.PublishAsync(json);

await publisher.ShutdownAsync(TimeSpan.FromSeconds(15));
