import { ObjectId, Timestamp } from "mongodb";

export function serializeMongoData(data: any): any {
  if (data === null || data === undefined) return data;
  
  if (Array.isArray(data)) {
    return data.map(item => serializeMongoData(item));
  }

  if (typeof data === "object") {
    const result: any = {};
    Object.entries(data).forEach(([key, value]) => {
      if (value instanceof ObjectId) {
        result[key] = value.toString();
      } else if (value instanceof Timestamp) {
        result[key] = new Date(value.getHighBits() * 1000).toISOString();
      } else if (value instanceof Date) {
        result[key] = value.toISOString();
      } else if (typeof value === "object" && value !== null) {
        result[key] = serializeMongoData(value);
      } else if (Array.isArray(value)) {
        result[key] = serializeMongoData(value);
      } else {
        // Convert any non-plain object to its string representation
        result[key] = typeof value === 'object' ? JSON.stringify(value) : value;
      }
    });
    return result;
  }
  return data;
}
