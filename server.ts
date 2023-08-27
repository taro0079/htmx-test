import { serve } from "https://deno.land/std@0.140.0/http/server.ts";

const HTML = await Deno.readFile("./index.html");

const CLICK_ROUTE = new URLPattern({ pathname: "/click" });
const SAMPLE_TEXT = "<p>RESPONSE TEXT</p>";

function handler(req: Request): Response {
  const headers = new Headers();

  // CORSヘッダの設定
  headers.set("Access-Control-Allow-Origin", "*");
  headers.set("content-type", "text/html");

  // OPTIONSメソッドのリクエストの場合、他のCORS関連のヘッダを設定
  if (req.method === "OPTIONS") {
    headers.set("Access-Control-Allow-Methods", "POST, GET, OPTIONS");
    headers.set("Access-Control-Allow-Headers", "hx-current-url, hx-request");
    return new Response(null, {
      headers: headers,
      status: 204, // No Content
    });
  }

  if (CLICK_ROUTE.exec(req.url)) {
    return new Response(SAMPLE_TEXT, {
      headers,
    });
  }

  return new Response(HTML, {
    headers,
  });
}

serve(handler);
