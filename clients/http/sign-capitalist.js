export function signCapitalistRequest(request, crypto) {
  const timestamp = Date.now().toString();
  const substitutedBody = request.body?.tryGetSubstituted?.();
  const body = substitutedBody == null ? "" : substitutedBody;
  const apiSecret = request.environment.get("apiSecret");
  const signature = crypto.sha256()
    .updateWithText(timestamp + body + apiSecret)
    .digest()
    .toHex();

  request.variables.set("timestamp", timestamp);
  request.variables.set("signature", signature);
}
