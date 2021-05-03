exports.handler = async function (event, context) {
    return {
        'statusCode': 200,
        'headers': {
            'Content-Type': 'text/plain'
        },
        'isBase64Encoded': false,
        'body': `${(Math.floor(Date.now() / 1e3)).toString()}`
    }
};