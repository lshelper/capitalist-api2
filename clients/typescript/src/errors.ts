export class CapitalistApi2Error extends Error {
  readonly status: number;
  readonly responseBody: string;

  constructor(message: string, status: number, responseBody: string) {
    super(message);
    this.name = 'CapitalistApi2Error';
    this.status = status;
    this.responseBody = responseBody;
  }
}
