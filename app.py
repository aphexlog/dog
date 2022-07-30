# create a lambda function


def handler(event, context):
    print(event)
    return event['message']
