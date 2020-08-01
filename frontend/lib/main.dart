// ignore: avoid_web_libraries_in_flutter
import 'dart:convert';
import 'dart:html' as html;

import 'package:auto_size_text/auto_size_text.dart';
import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;
import 'package:influ_dojo/model/rank_user.dart';
import 'package:influ_dojo/model/ranking.dart';
import 'package:influ_dojo/model/ranking_all.dart';

Future<RankingAll> fetchRanking() async {
  final response = await http.get('http://localhost:8080/api/ranking/all');

  if (response.statusCode == 200) {
    // If the server did return a 200 OK response,
    // then parse the JSON.
    return RankingAll.fromJson(json.decode(response.body));
  } else {
    // If the server did not return a 200 OK response,
    // then throw an exception.
    throw Exception('Failed to load all ranking');
  }
}

void main() {
  runApp(MyApp());
}

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      debugShowCheckedModeBanner: false,
      title: 'Influ Dojo',
      theme: ThemeData(
        primarySwatch: Colors.blue,
        visualDensity: VisualDensity.adaptivePlatformDensity,
      ),
      home: TabPage(),
    );
  }
}

class TabPage extends StatefulWidget {
  @override
  _TabPageState createState() => _TabPageState();
}

class _TabPageState extends State<TabPage> {
  Future<RankingAll> futureRanking;

  @override
  void initState() {
    super.initState();
    futureRanking = fetchRanking();
  }

  @override
  Widget build(BuildContext context) {
    return DefaultTabController(
      length: 3,
      child: Scaffold(
        appBar: AppBar(
          title: Text("駆け出しインフルエンサー道場"),
          bottom: TabBar(
            tabs: [
              Tab(text: "Today"),
              Tab(text: "Week"),
              Tab(text: "Month"),
            ],
          ),
        ),
        body: FutureBuilder<RankingAll>(
          future: futureRanking,
          builder: (context, snapshot) {
            if (snapshot.hasData) {
              return TabBarViewBuilder(snapshot.data);
            } else if (snapshot.hasError) {
              return Text("エラーが発生しました。お問い合わせ下さい。");
            }
            return CircularProgressIndicator();
          },
        ),
      ),
    );
  }
}

// ignore: non_constant_identifier_names
Widget TabBarViewBuilder(RankingAll rankingAll) {
  return TabBarView(
    children: [
      MyHomePage(rankingAll.dailyWorkRanking, rankingAll.dailyResultRanking),
      MyHomePage(rankingAll.weeklyWorkRanking, rankingAll.weeklyResultRanking),
      MyHomePage(rankingAll.monthlyWorkRanking, rankingAll.monthlyResultRanking),
    ],
  );
}

class MyHomePage extends StatefulWidget {
  final Ranking workRanking;
  final Ranking resultRanking;

  MyHomePage(this.workRanking, this.resultRanking);

  @override
  _MyHomePageState createState() => _MyHomePageState();
}

class _MyHomePageState extends State<MyHomePage> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: SingleChildScrollView(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: <Widget>[
            SizedBox(
              height: 16,
            ),
            Text(
              "現在の参加者: ${widget.workRanking.rankUsers.length}名",
              style: TextStyle(fontSize: 20),
            ),
            Padding(
              padding: const EdgeInsets.fromLTRB(8, 8, 8, 0),
              child: Card(
                child: Column(
                  children: [
                    Padding(
                      padding: const EdgeInsets.all(8),
                      child: Text(
                        '作業ランキング',
                        style: TextStyle(fontSize: 20),
                      ),
                    ),
                    for (int index = 0; index <= 9; index++)
                      buildCard(context, widget.workRanking.rankUsers[index]),
                  ],
                ),
              ),
            ),
            Padding(
              padding: const EdgeInsets.all(8.0),
              child: Card(
                child: Column(
                  children: [
                    Padding(
                      padding: const EdgeInsets.all(8),
                      child: Text(
                        '成果ランキング',
                        style: TextStyle(fontSize: 20),
                      ),
                    ),
                    for (int index = 0; index <= 9; index++)
                      buildCard(context, widget.resultRanking.rankUsers[index]),
                  ],
                ),
              ),
            ),
          ],
        ),
      ),
    );
  }

  Widget buildCard(BuildContext context, RankUser rankUser) {
    return ListTile(
      leading: Container(
        width: 100,
        child: Row(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            Text(
              "${rankUser.ranking}",
              style: TextStyle(fontSize: 18),
            ),
            SizedBox(width: 4),
            ChangeRank(rankUser.ranking, rankUser.lastRanking),
            SizedBox(width: 10),
            Image.network(rankUser.profileImage),
          ],
        ),
      ),
      title: AutoSizeText(
        rankUser.name,
        maxLines: 1,
        overflow: TextOverflow.ellipsis,
      ),
      subtitle: AutoSizeText(
        "@${rankUser.screenName}",
        maxLines: 1,
        overflow: TextOverflow.ellipsis,
      ),
      trailing: Text(
        "${rankUser.point} pt",
      ),
      onTap: () => _moveTwitter(rankUser.screenName),
    );
  }

  Widget ChangeRank(int ranking, int lastRanking) {
    String text = "";
    Color color;
    if (lastRanking == 0) {
      text = "N";
      color = Colors.orange;
    } else if (ranking < lastRanking) {
      text = "▲";
      color = Colors.blue;
    } else if (ranking > lastRanking) {
      text = "▼";
      color = Colors.red;
    } else if (ranking == lastRanking) {
      text = "-";
      color = Colors.black;
    }

    return Text(
      text,
      style: TextStyle(fontSize: 16, color: color),
    );
  }

  _moveTwitter(String screenName) async {
    print("tap!!");
    var url = "https://twitter.com/$screenName";
    html.window.open(url, "");
  }
}
